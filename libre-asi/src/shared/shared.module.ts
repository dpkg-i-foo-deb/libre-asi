import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { MatIconRegistry } from '@angular/material/icon';
import { DomSanitizer } from '@angular/platform-browser';

@NgModule({
  declarations: [],
  imports: [CommonModule, HttpClientModule],
  exports: [],
})
export class SharedModule {
  constructor(
    private registry: MatIconRegistry,
    private sanitizer: DomSanitizer
  ) {
    this.registry.addSvgIcon(
      'eye_on',
      this.sanitizer.bypassSecurityTrustResourceUrl('../../assets/eye_on.svg')
    );
    this.registry.addSvgIcon(
      'eye_off',
      this.sanitizer.bypassSecurityTrustResourceUrl('../../assets/eye_off.svg')
    );
  }
}
