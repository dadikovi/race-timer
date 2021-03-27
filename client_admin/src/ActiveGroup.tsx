import { Card, CardContent, Typography, Box, IconButton } from "@material-ui/core";
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

export default function ActiveGroup(props: ActiveGroupProps) {
    const throwError = useAsyncError();

    if (!props.data || props.data.group.id < 1) {
        return null
    }

    const groupStatus = groupStarted(props.data) ? `Started at ${new Date(props.data.group.start).toLocaleDateString()}`
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