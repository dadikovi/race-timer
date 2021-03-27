import React from 'react'
import {RaceResultsDo, SegmentDto} from './model'

export const useAsyncError = () => {
    const [_, setError] = React.useState();
    return React.useCallback(
      e => {
        setError(() => {
          throw e;
        });
      },
      [setError],
    );
  };

async function  execute<RESPONSE>(url: string, method?: string, body?: any): Promise<RESPONSE> {

    let init: RequestInit = { method : method }

    if (body) {
        init.body = JSON.stringify(body)
    }
    
    const response = await fetch(`http://localhost:8010${url}`, init)

    if (response.status === 200) {
        return response.json()
    } else {
        throw new Error(`Error while fetching ${url}: ${(await response.text())}`)
    }
}

function post<RESPONSE>(url:string, body?: any): Promise<RESPONSE> {
    return execute(url, 'POST', body)
}

function get<RESPONSE>(url:string): Promise<RESPONSE> {
    return execute(url);
}

export async function startActiveGroup() {
    return post('/groups/active')
}

export async function createSegment(segmentName: string | undefined): Promise<void> {
    return post('/segments', {name: segmentName})
}

export async function registerParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
        return
    }
    return post('/participants', {startNumber: startNumber})
}

export async function finishParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
        return
    }
    return post(`/participants/${startNumber}`)
}

export async function createGroup(segmentId: Number | undefined) {
    return post('/groups', {segmentId: segmentId})
}

export async function getResults(): Promise<RaceResultsDo> {
    return get('/race/results')
}

export async function getSegments(): Promise<SegmentDto[]> {
    return get('/segments')
}