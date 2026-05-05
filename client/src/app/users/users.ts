import { HttpClient, HttpParams } from "@angular/common/http";
import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";

@Component({
  selector: 'app-users',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './users.html',
  styleUrl: './users.css'
})
export class UsersComponent {
  userName = '';
  message = '';
  isError = false;

  constructor(private http: HttpClient) {}

  createUser() {
    if (!this.userName.trim()) return;
    const body = new HttpParams().set('name', this.userName);
    this.http.post('http://localhost:8080/users', body.toString(), {
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      responseType: 'text'
    }).subscribe({
      next: res => {
        this.message = '✓ ' + res;
        this.isError = false;
        this.userName = '';
      },
      error: err => {
        this.message = '✗ ' + (err.error || 'Something went wrong');
        this.isError = true;
      }
    });
  }
}