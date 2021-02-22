import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Directory } from 'src/app/core/models/directory';
import { Todo } from 'src/app/core/models/todo';
import { GeneralResponse } from 'src/app/core/models/user';
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
export class UserDashboardComponent implements OnInit {

  dirs: Directory[] = [Directory.fromJson({ title: 'General', id: '0' })]

  constructor(public dialog: MatDialog,
    private dirService: DirectoryService,
    private notificationService: NotificationService) { }


  async ngOnInit() {
    try {
      const dirs = await this.dirService.getDirectories("0");
      if (Array.isArray(dirs)) {
        this.dirs.push(...dirs)
      } else if (dirs != null) {
        throw new Error((dirs as GeneralResponse).message);
      }
    } catch (err) {
      this.notificationService.notify((err.error || err).message || 'Something went wrong!')
    }
  }

  async createDir() {
    const dialogRef = this.dialog.open(CreateDirectoryComponent, {
      width: '250px'
    });

    try {
      const dirForm = await dialogRef.afterClosed().toPromise()
      if ((dirForm || {}).name) {
        const newDir = Directory.fromJson({ title: dirForm.name, parentId: "0" })
        const dir = await this.dirService.createDirectory(newDir)
        if (typeof dir != "string") {
          this.dirs.push(dir)
        }
        this.notificationService.notify(`${dir.title} created`)
      }
    } catch (err) {
      this.notificationService.notify((err.error || err).message || 'Something went wrong!')
    }

  }

}
