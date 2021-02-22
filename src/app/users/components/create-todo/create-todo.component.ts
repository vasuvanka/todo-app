import { Component } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';

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

  constructor() { }

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

}
