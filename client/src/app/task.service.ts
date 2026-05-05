// task.service.ts — full corrected file

import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Task {
  ID: number;      // Go JSON uses capital keys by default
  Title: string;
  Done: boolean;
  UserID: number;
}

export interface User {
  ID: number;
  Name: string;
}

@Injectable({ providedIn: 'root' })
export class TaskService {
  private base = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getTasks(userId: number): Observable<Task[]> {
    return this.http.get<Task[]>(`${this.base}/tasks`, {
      params: new HttpParams().set('user', userId)
    });
  }

  getUsers(): Observable<User[]> {
  return this.http.get<User[]>(`${this.base}/users`);
  }

  createTask(title: string, userId: number): Observable<string> {
    const body = new HttpParams()
      .set('title', title)
      .set('user', userId);          // send userId to backend
    return this.http.post(`${this.base}/tasks`, body.toString(), {
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      responseType: 'text'
    });
  }

  markDone(id: number): Observable<string> {
    return this.http.get(`${this.base}/tasks/done`, {
      params: new HttpParams().set('id', id),
      responseType: 'text'
    });
  }
}