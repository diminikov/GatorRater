import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import { CoursesComponent } from './courses/courses.component';
import { CommonModule } from '@angular/common';
import { AboutUsComponent } from './about-us/about-us.component';
import { LoginComponent } from './login/login.component';
import{HttpClientModule} from '@angular/common/http';
import { SignupComponent } from './signup/signup.component'
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { DiscussionBoardComponent } from './discussion-board/discussion-board.component';




@NgModule({
  declarations: [
    HeaderComponent,
    AppComponent,
    CoursesComponent,
    AboutUsComponent,
    LoginComponent,
    SignupComponent,
    DiscussionBoardComponent,
  ],
  imports: [
    FormsModule,
    MatButtonModule,
    MatInputModule,
    BrowserModule,
    NoopAnimationsModule,
    HttpClientModule,
    ReactiveFormsModule,
    RouterModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

