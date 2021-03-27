import { Chip, Box } from "@material-ui/core";
import DoneIcon from '@material-ui/icons/Done';
import DirectionsRunIcon from '@material-ui/icons/DirectionsRun';
import { ParticipantDto } from './model';

interface ParticipantsProps {
    participants?: ParticipantDto[] | undefined
}

export default function Participants(props: ParticipantsProps) {    
    let participants = []
    if (props.participants) {
        for (let partipant of props.participants) {
        let icon = partipant.raceTimeMs > 0 ? <DoneIcon /> : <DirectionsRunIcon />
        let label = `#${partipant.startNumber}`

        if (partipant.raceTimeMs > 0) {
            label += ` (${partipant.raceTimeMs / 1000} s)`
        }
        participants.push(<Chip
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