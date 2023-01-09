import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SharedModule } from 'src/shared/shared.module';
import { LoginModule } from 'src/login/login.module';

@NgModule({
  declarations: [],
  imports: [CommonModule, SharedModule, LoginModule],
  exports: [],
})
export class CoreModule {}
