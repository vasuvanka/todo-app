import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from '../guards/auth.guard';
import { CreateDirectoryComponent } from './components/create-directory/create-directory.component';
import { CreateTodoComponent } from './components/create-todo/create-todo.component';
import { ProfileComponent } from './components/profile/profile.component';
import { RootDirectoryComponent } from './components/root-directory/root-directory.component';
import { UserDashboardComponent } from './components/user-dashboard/user-dashboard.component';
import { ViewDirectoryComponent } from './components/view-directory/view-directory.component';


const routes: Routes = [
  {
    path: 'dir',
    canActivate: [AuthGuard],
    component: UserDashboardComponent,
    children: [
      {
        path: '',
        component: RootDirectoryComponent
      },
      {
        path: ':id',
        component: ViewDirectoryComponent
      },
      {
        path: ':id/todos',
        component: CreateTodoComponent
      }
    ]
  },
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [AuthGuard],
  },
  {
    path: '**',
    redirectTo: 'dir'
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
