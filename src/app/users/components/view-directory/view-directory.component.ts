import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ActivatedRouteSnapshot, Router } from '@angular/router';
import { Todo } from 'src/app/core/models/todo';
import { GeneralResponse } from 'src/app/core/models/user';
import { NotificationService } from 'src/app/core/services/notification.service';
import { TodoService } from '../../services/todo.service';

@Component({
  selector: 'app-view-directory',
  templateUrl: './view-directory.component.html',
  styleUrls: ['./view-directory.component.css']
})
export class ViewDirectoryComponent implements OnInit {
  todos: Todo[] = []
  constructor(private router: Router,
    private todoService: TodoService,
    private notificationService: NotificationService,
    private activatedRoute: ActivatedRoute) { }

  async ngOnInit() {
    try {
      const todos = await this.todoService.getTodos(this.activatedRoute.snapshot.params['id'])
      if (Array.isArray(todos)) {
        this.todos = todos
      } else {
        throw new Error((todos as GeneralResponse).message);
      }
    } catch (err) {
      this.notificationService.notify(err.message)
    }
  }

}
