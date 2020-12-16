// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  
  //openWeatherAPI
  apiKey: '7c8174fd732efc1fbdbf1eef9905883a', // https://home.openweathermap.org/api_keys
  apiUrl: 'http://api.openweathermap.org/data/2.5',

  //my-budget-api
  budgetAPIURL: 'http://localhost:3000',

  //my-calendar-api
  calendarAPIURL: 'http://localhost:3001',

  //my-entertainment-api
  entertainmentAPIURL: 'http://localhost:3002',

  //my-news-api
  newsAPIURL: 'http://localhost:3003'
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
import 'zone.js/dist/zone-error';  // Included with Angular CLI.
