import { AppModule} from './app/app.module'
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
// import 'bootstrap/dist/js/bootstrap.bundle.min.js';


platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));



