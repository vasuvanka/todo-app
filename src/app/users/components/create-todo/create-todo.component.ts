import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Todo } from 'src/app/core/models/todo';
import { NotificationService } from 'src/app/core/services/notification.service';
import { TodoService } from '../../services/todo.service';

@Component({
  selector: 'app-create-todo',
  templateUrl: './create-todo.component.html',
  styleUrls: ['./create-todo.component.css']
})
export class CreateTodoComponent {
  todoForm = new FormGroup({
    title: new FormControl('', [Validators.required, Validators.minLength(2)]),
    description: new FormControl('', [Validators.required, Validators.minLength(4)]),
    dueDate: new FormControl('', [Validators.required]),
    priority: new FormControl('', [Validators.required])
  })
  minDate = new Date()

  constructor(private todoService: TodoService,
    private notificationService: NotificationService,
    private activatedRoute: ActivatedRoute) { }

  get title(): AbstractControl {
    return this.todoForm.get('title')
  }
  get description(): AbstractControl {
    return this.todoForm.get('description')
  }
  get dueDate(): AbstractControl {
    return this.todoForm.get('dueDate')
  }
  get priority(): AbstractControl {
    return this.todoForm.get('priority')
  }

  async onCreate() {
    try {
      const dirId = this.activatedRoute.snapshot.params['id'] || "0"
      const newTodo = Todo.fromJson(this.todoForm.value)
      newTodo.dirId = dirId
      const todo = await this.todoService.createTodo(newTodo)
      this.notificationService.notify(`${todo.title} created`)
    } catch (err) {
      this.notificationService.notify(err.message)
    }
  }

}
