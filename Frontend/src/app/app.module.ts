import { NgModule } from '@angular/core';
// import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { provideHttpClient } from '@angular/common/http';


@NgModule({
  declarations: [

  ],  
  imports: [
    // NgbModule,
    BrowserModule,
    AppRoutingModule,
    AppComponent
    
    
    
  ],
  providers: [    provideHttpClient(), 
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
