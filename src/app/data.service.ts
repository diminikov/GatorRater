import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface MyClassData {
  [classId: string]: {
    name: string;
    user: string [];
    posts: string[];
  };
}

@Injectable({
  providedIn: 'root'
})
export class DataService {
  currentUser: string;
  constructor(private http: HttpClient) {
    this.currentUser = '';
   }
   classData: MyClassData = {
    CEN3031: {
      name: 'Introduction to Software Engineering',
      user: [],
      posts: []
    },
    COP3503C: {
      name: 'Programming Fundamentals 2',
      user: [],
      posts: []
    },
    COP3502C: {
      name: 'Programming Fundamentals 1',
      user: [],
      posts: []
    },
    CDA3101: {
      name: 'Introduction to Computer Organization',
      user: [],
      posts: []
    }, 
    COP3530: {
      name: 'Data Structures and Algorithm',
      user: [],
      posts: []
    },
    COT3100: {
      name: 'Applications of Discrete Structures',
      user: [],
      posts: []
    },
    CIS4301: {
      name: 'Information and Database Systems 1',
      user: [],
      posts: []
    },
    CIS4914: {
      name: 'Senior Project',
      user: [],
      posts: []
    },
    CNT4007: {
      name: 'Computer Network Fundamentals',
      user: [],
      posts: []
    },
    COP4020: {
      name: 'Programming Language Concepts',
      user: [],
      posts: []
    },
    COP4533: {
      name: 'Algorithm Abstraction and Design',
      user: [],
      posts: []
    },
    COP4600: {
      name: 'Operating Systems',
      user: [],
      posts: []
    },
    EGS4034: {
      name: 'Engineering Ethics and Professionalism',
      user: [],
      posts: []
    }
  };

  ngOnInit(): void{
    this.getCEN3031Post(this.classData);
  }

  addPost(classId: string, post: string, user: string) {
    const url = 'http://localhost:8080/post';
    const data = 
    { class: classId,
      username: user, 
      body: post};
    return this.http.post(url, data);
  }

  helperget(classId: string){
    if (classId == 'CEN3031'){
      this.getCEN3031Post(this.classData);
    }else if(classId == 'EGS4034'){
      this.getEGS4034Post(this.classData);

    }else if(classId = 'COP3503C'){
      this.getCOP3503CPost(this.classData);
    }else if(classId = 'COP3502C'){
      this.getCOP3502CPost(this.classData);
    }else if(classId = 'CDA3101'){
      this.getCDA3101Post(this.classData);
    }
    else if(classId = 'COP3530'){
      this.getCOP3530Post(this.classData);
    }
    else if(classId = 'COT3100'){
      this.getCOT3100Post(this.classData);
    }
    else if(classId = 'CIS4301'){
      this.getCIS4301Post(this.classData);
    }
    else if(classId = 'CIS4914'){
      this.getCIS4914Post(this.classData);
    }
    else if(classId = 'CNT4007'){
      this.getCNT4007Post(this.classData);
    }
    else if(classId = 'COP4020'){
      this.getCOP4020Post(this.classData);
    }
    else if(classId = 'COP4533'){
      this.getCOP4533Post(this.classData);
    }
    else if(classId = 'COP4600'){
      this.getCOP4600Post(this.classData);
    }
    else if(classId = 'EGS4034'){
      this.getEGS4034Post(this.classData);
    }
    
  }

  getCEN3031Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/CEN3031';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['CEN3031'].user.push(username);
          classData['CEN3031'].posts.push(body);
          console.log(classData['CEN3031'].posts);
        });
      });
  }

  getCOP3503CPost(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP3503C';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP3503C'].user.push(username);
          classData['COP3503C'].posts.push(body);
          console.log(classData['COP3503C'].posts);
        });
      });
  }
  

  getCOP3502CPost(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP3502C';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP3502C'].user.push(username);
          classData['COP3502C'].posts.push(body);
          console.log(classData['COP3502C'].posts);
        });
      });
  }
  
  getCDA3101Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/CDA3101';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['CDA3101'].user.push(username);
          classData['CDA3101'].posts.push(body);
          console.log(classData['CDA3101'].posts);
        });
      });
  }
  
  getCOP3530Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP3530';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP3530'].user.push(username);
          classData['COP3530'].posts.push(body);
          console.log(classData['COP3530'].posts);
        });
      });
  }
  
  getCOT3100Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COT3100';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COT3100'].user.push(username);
          classData['COT3100'].posts.push(body);
          console.log(classData['COT3100'].posts);
        });
      });
  }
  
  getCIS4301Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/CIS4301';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['CIS4301'].user.push(username);
          classData['CIS4301'].posts.push(body);
          console.log(classData['CIS4301'].posts);
        });
      });
  }
  
  getCIS4914Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/CIS4914';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['CIS4914'].user.push(username);
          classData['CIS4914'].posts.push(body);
          console.log(classData['CIS4914'].posts);
        });
      });
  }
  
  getCNT4007Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/CNT4007';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['CNT4007'].user.push(username);
          classData['CNT4007'].posts.push(body);
          console.log(classData['CNT4007'].posts);
        });
      });
  }
  
  getCOP4020Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP4020';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP4020'].user.push(username);
          classData['COP4020'].posts.push(body);
          console.log(classData['COP4020'].posts);
        });
      });
  }

  getCOP4533Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP4533';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP4533'].user.push(username);
          classData['COP4533'].posts.push(body);
          console.log(classData['COP4533'].posts);
        });
      });
  }

  getCOP4600Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/COP4600';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['COP4600'].user.push(username);
          classData['COP4600'].posts.push(body);
          console.log(classData['COP4600'].posts);
        });
      });
  }


  getEGS4034Post(classData: MyClassData) {
    const url = 'http://localhost:8080/post/class/EGS4034';
    this.http.get<any[]>(url)
      .subscribe(data => {
        // Loop through the data and add the username and body to their respective arrays in the CEN3031 object
        data.forEach(item => {
          const username = item.Username;
          const body = item.Body;
          classData['EGS4034'].user.push(username);
          classData['EGS4034'].posts.push(body);
          console.log(classData['EGS4034'].posts);
        });
      });
  }

  
  



  getClassData(classId: string) {
    return this.classData[classId];
  }

  /*addPost(classId: string, post: string) {
    this.classData[classId].posts.push(post);
  }
  */

  getClass(classId: string) {
    return this.classData[classId];
  }



  setCurrentUser(user: string) {
    this.currentUser = user;
  }

  getCurrentUser() {
    return this.currentUser;
  }

  logout() {
    this.currentUser = '';
  }






}



