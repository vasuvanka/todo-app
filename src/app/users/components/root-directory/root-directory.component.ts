import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Directory } from 'src/app/core/models/directory';
import { DirectoryService } from 'src/app/core/services/directory.service';
import { NotificationService } from 'src/app/core/services/notification.service';

@Component({
  selector: 'app-root-directory',
  templateUrl: './root-directory.component.html',
  styleUrls: ['./root-directory.component.css']
})
export class RootDirectoryComponent implements OnInit {

  dirs: Directory[] = []
  constructor(private dirService: DirectoryService,
    private notificationService: NotificationService) { }

  async ngOnInit() {
    try {
      const dirs = await this.dirService.getDirectories("0")
      if (Array.isArray(dirs)) {
        this.dirs = dirs
      } else {
        throw new Error((dirs as any).message)
      }
    } catch (err) {
      this.notificationService.notify(err.message)
    }
  }

  trackBy(dir: Directory): string {
    return dir.id
  }

}
