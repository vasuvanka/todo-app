import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-share-todo',
  templateUrl: './share-todo.component.html',
  styleUrls: ['./share-todo.component.css']
})
export class ShareTodoComponent {

  shareForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email])
  })
  constructor(
    public dialogRef: MatDialogRef<ShareTodoComponent>) { }

  get email(): AbstractControl {
    return this.shareForm.get('email')
  }

  onShare() {
    this.dialogRef.close(this.shareForm.value)
  }

}
