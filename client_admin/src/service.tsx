async function  execute<RESPONSE>(url: string, method?: string, body?: any): Promise<RESPONSE> {

    let init: RequestInit = { method : method }

    if (body) {
        init.body = JSON.stringify(body)
    }

    return (await fetch(`http://localhost:8010${url}`, init)).json()
}

function post<RESPONSE>(url:string, body?: any): Promise<RESPONSE> {
    return execute(url, 'POST', body)
}

function get<RESPONSE>(url:string): Promise<RESPONSE> {
    return execute(url);
}

export function startActiveGroup() {
    post('/groups/active')
}

export function createSegment(segmentName: string | undefined, onSegmentCreated: Function | undefined) {
    fetch("http://localhost:8010/segments", {
      method: 'POST',
      body: JSON.stringify({name: segmentName})
    }) 
      .then(res => res.json())
      .then(
        () => {
          if (onSegmentCreated !== undefined) {
            onSegmentCreated()
          }
        },
        (error) => {
        }
      )
}

export function registerParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
    return
    }
    fetch("http://localhost:8010/participants", {
    method: 'POST',
    body: JSON.stringify({startNumber: startNumber})
    });
}

export function finishParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
        return
    }
    fetch(`http://localhost:8010/participants/${startNumber}`, {
        method: 'POST'
    });
}

export function createGroup(segmentId: Number | undefined) {
    fetch("http://localhost:8010/groups", {
        method: 'POST',
        body: JSON.stringify({segmentId: segmentId})
    });
}

export function getResults(setResults: Function) {
  fetch("http://localhost:8010/race/results") 
  .then(res => res.json())
  .then(
    (result) => { setResults(result) }
  )
}

export function getSegments(setSegments: Function) {
  fetch("http://localhost:8010/segments") 
    .then(res => res.json())
    .then(
      (result) => { setSegments(result) }
    )
}