import { Component } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-create-directory',
  templateUrl: './create-directory.component.html',
  styleUrls: ['./create-directory.component.css']
})
export class CreateDirectoryComponent {

  directoryForm = new FormGroup({
    name: new FormControl('', [Validators.required])
  })
  constructor(
    public dialogRef: MatDialogRef<CreateDirectoryComponent>) { }

  get name(): AbstractControl {
    return this.directoryForm.get('name')
  }

  onCreate() {
    this.dialogRef.close(this.directoryForm.value)
  }

}
