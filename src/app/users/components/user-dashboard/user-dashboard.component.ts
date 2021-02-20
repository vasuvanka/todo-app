import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Directory } from 'src/app/core/models/directory';
import { Todo } from 'src/app/core/models/todo';
import { DirectoryService } from 'src/app/core/services/directory.service';
import { NotificationService } from 'src/app/core/services/notification.service';
import { TodoService } from '../../services/todo.service';
import { CreateDirectoryComponent } from '../create-directory/create-directory.component';
import { CreateTodoComponent } from '../create-todo/create-todo.component';

@Component({
  selector: 'app-user-dashboard',
  templateUrl: './user-dashboard.component.html',
  styleUrls: ['./user-dashboard.component.css']
})
export class UserDashboardComponent {

  dirs: Directory[] = []
  constructor(public dialog: MatDialog,
    private dirService: DirectoryService,
    private notificationService: NotificationService,
    private router: Router,
    private route: ActivatedRoute) { }


  async createDir() {
    const dialogRef = this.dialog.open(CreateDirectoryComponent, {
      width: '250px'
    });

    try {
      const dirForm = await dialogRef.afterClosed().toPromise()
      if ((dirForm || {}).name) {
        const parentId = this.route.snapshot.params['id'] || "0"
        const newDir = Directory.fromJson({ title: dirForm.name, parentId })
        const dir = await this.dirService.createDirectory(newDir)
        this.notificationService.notify(`${dir.title} created`)
      }
    } catch (err) {
      this.notificationService.notify(err.message)
    }

  }

  redirect() {
    this.router.navigate([this.route.snapshot.params['id'] || "0", 'todos'], {
      relativeTo: this.route
    })
  }

}
