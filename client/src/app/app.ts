// app.ts
import { Component } from '@angular/core';
import { RouterOutlet, RouterLink, RouterLinkActive } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, RouterLinkActive],
  template: `
    <nav>
      <a class="brand" routerLink="/">TaskFlow</a>
      <a routerLink="/tasks" routerLinkActive="active">Tasks</a>
      <a routerLink="/users" routerLinkActive="active">Users</a>
    </nav>
    <router-outlet />
  `,
  styles: []
})
export class App {}