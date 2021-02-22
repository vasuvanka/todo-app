import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { Todo } from 'src/app/core/models/todo';
import { GeneralResponse } from 'src/app/core/models/user';
import { NotificationService } from 'src/app/core/services/notification.service';
import { TodoService } from '../../services/todo.service';
import { CreateTodoComponent } from '../create-todo/create-todo.component';
import { ShareTodoComponent } from '../share-todo/share-todo.component';

@Component({
  selector: 'app-view-directory',
  templateUrl: './view-directory.component.html',
  styleUrls: ['./view-directory.component.css']
})
export class ViewDirectoryComponent implements OnInit {
  todos: Todo[] = []
  statusForm = new FormGroup({
    status: new FormControl(''),
    priority: new FormControl('')
  })
  constructor(private todoService: TodoService,
    private notificationService: NotificationService,
    readonly dialog: MatDialog,
    private activatedRoute: ActivatedRoute) { }

  async ngOnInit() {
    this.activatedRoute.params.subscribe(async params => {
      try {
        this.todos = []
        const todos = await this.todoService.getTodos(params['id'])
        if (Array.isArray(todos)) {
          this.todos = todos
        } else if (todos != null) {
          throw new Error((todos as GeneralResponse).message);
        }
      } catch (err) {
        this.notificationService.notify((err.error || err).message || 'Something went wrong!')
      }
    })

    this.statusForm.valueChanges.subscribe(form => {
      let index;
      if (form.status) {
        const [id, status] = form.status.split('.')
        index = this.todos.findIndex(t => t.id == id)
        this.todos[index].status = status
      }
      if (form.priority) {
        const [id, priority] = form.priority.split('.')
        index = index ? index : this.todos.findIndex(t => t.id == id)
        this.todos[index].priority = priority
      }
      this.todoService.updateTodo(this.todos[index])
    })
  }

  async createTodo() {
    const dialogRef = this.dialog.open(CreateTodoComponent, {
      width: '350px'
    });

    try {
      const todoForm = await dialogRef.afterClosed().toPromise()
      if (todoForm) {
        const todoObj = Todo.fromJson(todoForm)
        todoObj.dirId = this.activatedRoute.snapshot.params['id']
        const todo = await this.todoService.createTodo(todoObj)
        if (typeof todo != "string") {
          this.todos.unshift(todo)
          this.notificationService.notify(`${todo.title} created`)
        }
      }
    } catch (err) {
      this.notificationService.notify((err.error || err).message || 'Something went wrong!')
    }

  }

  async remove(todo: Todo) {
    try {
      const response = await this.todoService.removeTodo(todo.id)
      if (response.message == "removed") {
        this.notificationService.notify(`${todo.title} removed.`)
        const i = this.todos.findIndex(t => t.id == todo.id)
        this.todos.splice(i, 1)
      } else {
        throw new Error(response.message);
      }
    } catch (err) {
      this.notificationService.notify((err.error || err).message || 'Something went wrong!')
    }
  }

  getColor(todo: Todo): string {
    let color = 'default'
    switch (todo.status) {
      case 'inprogress':
        color = 'accent'
        break;
      case 'completed':
        color = 'primary'
        break;
    }
    return color;
  }

  getPriority(todo: Todo): string {
    let color = 'default'
    if (todo.priority == "medium") {
      color = "accent"
    } else if (todo.priority == "high") {
      color = "warn"
    }
    return color;
  }

  async share(todo: Todo) {
    const dialogRef = this.dialog.open(ShareTodoComponent, {
      width: '350px'
    });

    try {
      const shareForm = await dialogRef.afterClosed().toPromise()
      if (shareForm) {
        const res = await this.todoService.shareTodo(todo.id, shareForm.email)
        this.notificationService.notify(`${todo.title} ${res.message}`)
      }
    } catch (err) {
      this.notificationService.notify((err.error || err).message || 'Something went wrong!')
    }
  }

}
