import { Component, OnInit} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';



interface user{
  id: string;
  username: string;
  password: string;
}


interface someClass {
  name: string;
  professor: string;
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


@Injectable({ 
  providedIn: 'root' 
 }) 


export class AppComponent{
 

  constructor(private httpClient: HttpClient, router:Router){

  }

  

  
  
  public users: user[] = [];



loadedFeature = 'course';

onNavigate(feature: string){
  this.loadedFeature = feature;
}

  title = 'GatorRater';

  /*
  fetchUsers(){
    this.httpClient.get('/http://localhost:8080/api/users').subscribe(
      users => {
        console.log(users);
      }
    );
    console.log('Fetch1');
  }
  */

  /*
  public getUsers(): Observable<any> {
    const url = `/http://localhost:8080/api/users`;
    return this.httpClient.get<any>(url);
    console.log('Fetch2');
}
*/

private baseUrl = 'http://localhost:8080/api/users';

  getUser(){
    this.httpClient.get<user[]>(this.baseUrl).subscribe( data=>{
      this.users = data;
      console.log(data);
    });
  
    console.log('Fetch3');
  }


  RunTest(){
    this.getUser();
    //this.fetchUsers();
    //this.getUsers(); 
  }






}
