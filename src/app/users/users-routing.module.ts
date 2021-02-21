import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from '../guards/auth.guard';
import { DefaultDirectoryComponent } from './components/default-directory/default-directory.component';
import { ProfileComponent } from './components/profile/profile.component';
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
        component: DefaultDirectoryComponent
      },
      {
        path: ':id',
        component: ViewDirectoryComponent
      }
    ]
  },
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [AuthGuard],
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
