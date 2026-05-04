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
  userId = 1; // hardcoded for now

  // inject the service — Angular creates it for us
  constructor(private taskService: TaskService) {}

  // runs once when component loads
  ngOnInit() {
    this.loadTasks();
  }

  loadTasks() {
    this.taskService
      .getTasks(this.userId)
      .subscribe(tasks => this.tasks = tasks);
  }

  addTask() {
    if (!this.newTitle.trim()) return;
    this.taskService
      .createTask(this.newTitle)
      .subscribe(() => {
        this.newTitle = '';
        this.loadTasks(); // refresh list
      });
  }

  complete(task: Task) {
    this.taskService
      .markDone(task.id)
      .subscribe(() => task.done = true);
  }
}