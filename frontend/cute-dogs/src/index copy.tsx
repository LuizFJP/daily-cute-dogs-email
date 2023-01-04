import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './index.css';
import App from './home/App';
import reportWebVitals from './home/reportWebVitals';
import CancelSubscription from './cancelSubscription/CancelSubscription';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path: "/cancel-subscription",
    element: <CancelSubscription />
  }
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
)

  

// If you want to start measuring performance in your app, pass a function
// to lrouog results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
