import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { Task, User, TaskService } from '../task.service';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterLink],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css',
  encapsulation: ViewEncapsulation.None
})
export class TasksComponent implements OnInit {
  tasks: Task[] = [];
  users: User[] = [];
  newTitle = '';
  selectedUserId: number | null = null;

  constructor(private taskService: TaskService) {}

  ngOnInit() {
    this.taskService.getUsers().subscribe({
      next: users => this.users = users ?? [],
      error: err => console.error('Failed to load users', err)
    });
  }

  onUserSelect(userId: number) {
    this.selectedUserId = userId;
    this.tasks = [];
    this.loadTasks();
  }

  loadTasks() {
    if (!this.selectedUserId) return;
    this.taskService
      .getTasks(this.selectedUserId)
      .subscribe({
        next: tasks => this.tasks = tasks ?? [],
        error: err => console.error('Failed to load tasks', err)
      });
  }

  addTask() {
    if (!this.newTitle.trim() || !this.selectedUserId) return;
    this.taskService
      .createTask(this.newTitle, this.selectedUserId)
      .subscribe({
        next: () => {
          this.newTitle = '';
          this.loadTasks();
        },
        error: err => console.error('Failed to create task', err)
      });
  }

  complete(task: Task) {
    this.taskService
      .markDone(task.ID)
      .subscribe({
        next: () => task.Done = true,
        error: err => console.error('Failed to mark done', err)
      });
  }

  getSelectedUserName(): string {
    const user = this.users.find(u => u.ID === this.selectedUserId);
    return user ? user.Name : '';
  }
}