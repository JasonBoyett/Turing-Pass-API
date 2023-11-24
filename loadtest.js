// This load test uses the k6 testing framework
// Run: k6 run loadtest.js 
// get k6 from https://k6.io
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  // vus = virtual users
  vus: 5000,
  duration: '1s',
  ext: {
    loadimpact: {
      projectID: 3671030,
      name: "Load Test",
    },
  },
}

export default function () {
  http.get('https://turing-pass-api-production.up.railway.app/jsonp?passWord=foo&siteName=bar&symbols=true');
  sleep(1);
}
