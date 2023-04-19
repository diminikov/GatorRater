import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface MyClassData {
  [classId: string]: {
    name: string;
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
  private classData: MyClassData = {
    CEN3031: {
      name: 'Introduction to Software Engineering',
      posts: []
    },
    COP3503C: {
      name: 'Programming Fundamentals 2',
      posts: []
    },
    COP3502C: {
      name: 'Programming Fundamentals 1',
      posts: []
    },
    CDA3101: {
      name: 'Introduction to Computer Organization',
      posts: []
    }, 
    COP3530: {
      name: 'Data Structures and Algorithm',
      posts: []
    },
    COT3100: {
      name: 'Applications of Discrete Structures',
      posts: []
    },
    CIS4301: {
      name: 'Information and Database Systems 1',
      posts: []
    },
    CIS4914: {
      name: 'Senior Project',
      posts: []
    },
    CNT4007: {
      name: 'Computer Network Fundamentals',
      posts: []
    },
    COP4020: {
      name: 'Programming Language Concepts',
      posts: []
    },
    COP4533: {
      name: 'Algorithm Abstraction and Design',
      posts: []
    },
    COP4600: {
      name: 'Operating Systems',
      posts: []
    },
    EGS4034: {
      name: 'Engineering Ethics and Professionalism',
      posts: []
    }
  };

  addPost(classId: string, post: string) {
    const url = `/api/posts/${classId}`;
    return this.http.post(url, { post });
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



