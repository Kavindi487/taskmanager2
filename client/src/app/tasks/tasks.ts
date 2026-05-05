// tasks.ts — full corrected file

import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Task, TaskService } from '../task.service';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class TasksComponent implements OnInit {
  tasks: Task[] = [];
  newTitle = '';
  userId = 1;

  constructor(private taskService: TaskService) {}

  ngOnInit() {
    this.loadTasks();
  }

  loadTasks() {
    this.taskService
      .getTasks(this.userId)
      .subscribe(tasks => this.tasks = tasks ?? []);
  }

  addTask() {
    if (!this.newTitle.trim()) return;
    this.taskService
      .createTask(this.newTitle, this.userId)   // pass userId
      .subscribe(() => {
        this.newTitle = '';
        this.loadTasks();
      });
  }

  complete(task: Task) {
    this.taskService
      .markDone(task.ID)                        // use capital ID
      .subscribe(() => task.Done = true);       // use capital Done
  }
}