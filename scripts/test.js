import http from 'k6/http'

export let options = {
    vus:5,
    durations:'5s'
}

export default function(){
    http.get('http://host.docker.internal:8000/hello')
}