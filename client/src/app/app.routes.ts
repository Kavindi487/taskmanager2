import { Routes } from '@angular/router';
import { TasksComponent } from './tasks/tasks';
import { UsersComponent } from './users/users';

export const routes: Routes = [
  { path: '',       redirectTo: 'tasks', pathMatch: 'full' },
  { path: 'tasks',  component: TasksComponent },
  { path: 'users',  component: UsersComponent },
];