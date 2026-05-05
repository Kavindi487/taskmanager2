// app.ts — add RouterLink
import { Component } from '@angular/core';
import { RouterOutlet, RouterLink } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink],   // add RouterLink here
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {}