export interface SegmentDto {
    name: string;
    id: number;
}
  
export interface RaceResultsDo {
    activeGroup: ParticipantDto[];
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