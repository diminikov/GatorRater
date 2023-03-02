import {Component} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { OnInit } from '@angular/core';
/**
 * 
* @title login demo
*/
export interface  user{
    username: string;
    password: string;
}
@Component({
selector: 'app-login',
styleUrls: ['login.component.css'],
templateUrl: 'login.component.html',
})
export class LoginComponent{
constructor(
    private httpClient: HttpClient
) {}
    
username : string ="";
password : string ="";
show: boolean= false;
public users: user[] = [];
submit(){
console.log("user name is " + this.username)
this.clear();
}
clear(){
this.username ="";
this.password = "";
this.show = true;
}

ngOnInit(): void{
this.getAllUsers().subscribe(data=> this.users = data)
}
getAllUsers(): Observable<user[]>{
    return this.httpClient.get<user[]>("http/localhost:8080/users");    
}
}

