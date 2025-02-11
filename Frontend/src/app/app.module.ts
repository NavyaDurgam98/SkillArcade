import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { provideHttpClient } from '@angular/common/http';

@NgModule({
  declarations: [],  
  imports: [
    BrowserModule,
    AppRoutingModule,
    AppComponent
    
  ],
  providers: [    provideHttpClient(), 
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
