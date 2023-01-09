import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StatusbarComponent } from './statusbar/statusbar.component';



@NgModule({
  declarations: [
    StatusbarComponent
  ],
  imports: [
    CommonModule
  ],
  exports:[
    StatusbarComponent
  ]
})
export class CoreModule { }
