import { NgModule } from "@angular/core";
import { RouterModule,Routes } from "@angular/router";
import { AboutUsComponent } from "./about-us/about-us.component";
import { CoursesComponent } from "./courses/courses.component";
import { LoginComponent } from "./login/login.component";
import { SignupComponent } from "./signup/signup.component";
import { DiscussionBoardComponent } from "./discussion-board/discussion-board.component";
const routes: Routes = [
    {path: '', redirectTo: 'login', pathMatch: 'full'},
    {path: 'login', component: LoginComponent},
    {path: 'signup', component: SignupComponent},
    {path: 'aboutus', component: AboutUsComponent},
    {path: 'Courses', component: CoursesComponent},
    {path: 'DiscussionBoard', component: DiscussionBoardComponent}


];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})

export class AppRoutingModule {}