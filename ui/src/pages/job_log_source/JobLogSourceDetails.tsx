import { createStyles, Theme, WithStyles, Typography, TextField, Button } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";

import { IJobLogSourceStore, IJobLogSourceStoreState } from './JobLogSourceStore';
import { ResourceText } from '../../i18n/Languages';

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @JobLogSourceDetails }
 */
export interface IJobLogSourceDetailsProps
    extends WithStyles<typeof styles>, IUseStoreFromContainerContextProps
{
}

/**
 * Implementation of JobLogSourceDetails.
 */
const JobLogSourceDetailsImpl: React.FunctionComponent<IJobLogSourceDetailsProps> = props => {
    const { classes } = props;

    const [ storeState, store ] = useStoreStateFromContainerContext<IJobLogSourceStore, IJobLogSourceStoreState>(props);
    const setName = useCallback(e => store.setName(e.target.value), [ store ]);
    const setSourceFolder = useCallback(e => store.setSourceFolder(e.target.value), [ store ]);
    const setFileRegex = useCallback(e => store.setFileRegex(e.target.value), [ store ]);
    const setSuccessRegex = useCallback(e => store.setSuccessRegex(e.target.value), [ store ]);
    const setErrorRegex = useCallback(e => store.setErrorRegex(e.target.value), [ store ]);
    const save = useCallback(e => store.save(), [ store ]);

    if (!storeState)
    {
        return null;
    }

    return <div className={classes.root}>
        <TextField label={<ResourceText getText={x => x.log_source_name} />}
                    value={storeState.jobLogSource.name}
                    onChange={setName}></TextField>
        <TextField label={<ResourceText getText={x => x.log_source_typ} />}
                    value={storeState.jobLogSource.type}
                    disabled></TextField>
        <TextField label={<ResourceText getText={x => x.log_source_folder} />}
                    value={storeState.jobLogSource.sourceFolder}
                    onChange={setSourceFolder}></TextField>
        <TextField label={<ResourceText getText={x => x.log_source_file_regex} />}
                    value={storeState.jobLogSource.fileRegex}
                    onChange={setFileRegex}></TextField>
        <TextField label={<ResourceText getText={x => x.job_log_source_success_regex} />}
                    value={storeState.jobLogSource.successRegex}
                    onChange={setSuccessRegex}></TextField>
        <TextField label={<ResourceText getText={x => x.job_log_source_error_regex} />}
                    value={storeState.jobLogSource.errorRegex}
                    onChange={setErrorRegex}></TextField>
        <Button onClick={save}>
            <ResourceText getText={x => x.save}/>
        </Button>
    </div>;
};

/**
 * Component that shows all properties of a job log source.
 */
export const JobLogSourceDetails = withStyles(styles)(JobLogSourceDetailsImpl);
