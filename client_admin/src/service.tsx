import {RaceResultsDo, SegmentDto} from './model'

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

export async function createSegment(segmentName: string | undefined): Promise<void> {
    await post('/segments', {name: segmentName})
}

export function registerParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
        return
    }
    post('/participants', {startNumber: startNumber})
}

export function finishParticipant(startNumber: Number | undefined) {
    if (!startNumber) {
        return
    }
    post(`/participants/${startNumber}`)
}

export function createGroup(segmentId: Number | undefined) {
    post('/groups', {segmentId: segmentId})
}

export async function getResults(): Promise<RaceResultsDo> {
    const results: RaceResultsDo = await get('/race/results')
    return results
}

export async function getSegments(): Promise<SegmentDto[]> {
    const results: SegmentDto[] = await get('/segments')
    return results
}