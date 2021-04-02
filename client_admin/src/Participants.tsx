import { Chip, Box } from "@material-ui/core";
import DoneIcon from '@material-ui/icons/Done';
import DirectionsRunIcon from '@material-ui/icons/DirectionsRun';
import { ParticipantDto } from './model';

interface ParticipantsProps {
    participants?: ParticipantDto[] | undefined
}

export function msToHumanReadableDuration(ms: number) {
    let response: string = ""
    const seconds = Math.round(ms / 1000)
    const numyears = Math.floor(seconds / 31536000);
    const numdays = Math.floor((seconds % 31536000) / 86400); 
    const numhours = Math.floor(((seconds % 31536000) % 86400) / 3600);
    const numminutes = Math.floor((((seconds % 31536000) % 86400) % 3600) / 60);
    const numseconds = (((seconds % 31536000) % 86400) % 3600) % 60;

    if (numyears) {
        response += numyears + " years"
    }
    if (numdays) {
        if (response) {
            response += " "
        }
        response += numdays + " days"
    }
    if (numhours) {
        if (response) {
            response += " "
        }
        response += numhours + " hours"
    }
    if (numminutes) {
        if (response) {
            response += " "
        }
        response += numminutes + " m"
    }
    if (numseconds) {
        if (response) {
            response += " "
        }
        response += numseconds + " s"
    }

    return response;
}

export default function Participants(props: ParticipantsProps) {    
    let participants = []
    if (props.participants) {
        for (let partipant of props.participants) {
        let icon = partipant.raceTimeMs > 0 ? <DoneIcon /> : <DirectionsRunIcon />
        let label = `#${partipant.startNumber}`

        if (partipant.raceTimeMs > 0) {
            label += ` (${msToHumanReadableDuration(partipant.raceTimeMs)})`
        }
        participants.push(<Chip
            key={label}
            variant="outlined"
            icon={icon}
            label={label}
            color="default"/>
          )
        }
    }
    
    const emptyState = participants.length > 0 ? '' : 'No participants yet.'


    return (<Box className="formbox" display="flex" alignItems="center">
        {participants}
        {emptyState}
    </Box>)
}