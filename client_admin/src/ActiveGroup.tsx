import { Card, CardContent, Typography, Box, IconButton, Divider } from "@material-ui/core";
import PlayCircleFilledIcon from '@material-ui/icons/PlayCircleFilled';
import { ActiveGroupResultsDto } from './model';
import Participants from './Participants'
import { startActiveGroup, useAsyncError } from './service';

interface ActiveGroupProps {
    onRefresh: Function;
    data?: ActiveGroupResultsDto
}

function groupStarted(groupResults: ActiveGroupResultsDto): boolean {
    const startDate = new Date(groupResults.group.start);
    return startDate.getFullYear() > 1 // Go server represents empty date as 0001-01-01
}

function msToHumanReadableDuration(ms: number) {
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

export default function ActiveGroup(props: ActiveGroupProps) {
    const throwError = useAsyncError();

    if (!props.data || props.data.group.id < 1) {
        return null
    }

    const groupStatus = groupStarted(props.data) ? <div style={{padding: '0px 10px'}}>
            Started {msToHumanReadableDuration(new Date().getTime() - new Date(props.data.group.start).getTime())} ago
        </div> 
        : <IconButton size="medium" color="secondary" onClick={() => {
            startActiveGroup()
                .then(() => { if (props.onRefresh) {props.onRefresh()}})
                .catch((err) => throwError(err));
            }}>
            <PlayCircleFilledIcon />
        </IconButton>

    return (<Card elevation = {3}>
        <CardContent>
            <Box display="flex" alignItems="center">
                <Typography variant="h5" component="h2">Active group</Typography>
                {groupStatus}
            </Box>
            <Participants participants={props.data?.participants}></Participants>            
        </CardContent>
    </Card>)
}