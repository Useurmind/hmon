import { createStyles, Theme, WithStyles } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";

import { ILogSourcePageStore, ILogSourcePageStoreState } from "./LogSourcePageStore";
import { LogSourceList } from './LogSourceList';

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @LogSourcesPage }
 */
export interface ILogSourcesPageProps
    extends WithStyles<typeof styles>
{
}

/**
 * Implementation of LogSourcesPage.
 */
const LogSourcesPageImpl: React.FunctionComponent<ILogSourcesPageProps> = props => {
    const { classes } = props;

    // const [ storeState, store ] = useStoreStateFromContainerContext<ILogSourcePageStore, ILogSourcePageStoreState>(props);
    // // const  = useCallback(() => store..trigger(1), [ store ]);

    // if (!storeState)
    // {
    //     return null;
    // }

    return <div className={classes.root}>
        <LogSourceList storeRegistrationKey="ILogSourcePageStore"></LogSourceList>
    </div>;
};

/**
 * Component that shows the log sources.
 */
export const LogSourcesPage = withStyles(styles)(LogSourcesPageImpl);
