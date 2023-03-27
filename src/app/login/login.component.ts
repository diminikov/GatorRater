import {Component,OnInit} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Users } from '../users';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';




@Component({
selector: 'app-login',
styleUrls: ['login.component.css'],
templateUrl: 'login.component.html',
})


export class LoginComponent implements OnInit{

  public loginForm!: FormGroup
  constructor( private formBuilder : FormBuilder, private http : HttpClient, private router: Router) {}

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      email:[''],
      password:['']
    })
  }
  private baseUrl = 'http://localhost:8080/api/users';

  login(){
    this.http.get<any>(this.baseUrl)
    .subscribe(res=>{
      const user = res.find((a:any)=>{
        return a.email === this.loginForm.value.email && a.password === this.loginForm.value.password
      });
      if(user){
        alert("Login Success");
        this.loginForm.reset();
        this.router.navigate(['Courses'])
      }else{
        alert("user not found");
      }
    }, err=>{
      alert("Something went wrong")
    })

  }


}