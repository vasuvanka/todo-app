import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { ProfileComponent } from './components/profile/profile.component';
import { UserDashboardComponent } from './components/user-dashboard/user-dashboard.component';
import { MaterialModule } from '../material/material.module';
import { MatDialogModule } from '@angular/material/dialog';
import { CreateDirectoryComponent } from './components/create-directory/create-directory.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ViewDirectoryComponent } from './components/view-directory/view-directory.component';
import { RootDirectoryComponent } from './components/root-directory/root-directory.component';
import { CreateTodoComponent } from './components/create-todo/create-todo.component';
import { MatDatepickerModule } from '@angular/material/datepicker';

import { MatRadioModule } from '@angular/material/radio';
import { MatNativeDateModule } from '@angular/material/core';

@NgModule({
  declarations: [ProfileComponent,
    UserDashboardComponent, CreateDirectoryComponent,
    ViewDirectoryComponent, RootDirectoryComponent, CreateTodoComponent],
  imports: [
    CommonModule,
    UsersRoutingModule,
    MaterialModule,
    MatDialogModule,
    ReactiveFormsModule,
    MatNativeDateModule,
    MatDatepickerModule,
    MatRadioModule
  ]
})
export class UsersModule { }
