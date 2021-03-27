export interface SegmentDto {
    name: string;
    id: number;
}
 
export interface GroupDto {
    start: Date
    id: number
}

export interface ActiveGroupResultsDto {
    group: GroupDto
    participants: ParticipantDto[];
}

export interface RaceResultsDo {
    activeGroup: ActiveGroupResultsDto;
    segments: SegmentResultsDto[];
}

export interface SegmentResultsDto {
    segmentName: string;
    participants: ParticipantDto[];
}

export interface ParticipantDto {
    startNumber: number;
    groupId: number;
    raceTimeMs: number;
}