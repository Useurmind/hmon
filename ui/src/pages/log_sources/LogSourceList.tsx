import { createStyles, Theme, WithStyles, List, ListItem, ListItemText } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";

import { ILogSourcePageStore, ILogSourcePageStoreState } from "./LogSourcePageStore";

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @LogSourceList }
 */
export interface ILogSourceListProps
    extends WithStyles<typeof styles>, IUseStoreFromContainerContextProps
{
}

/**
 * Implementation of LogSourceList.
 */
const LogSourceListImpl: React.FunctionComponent<ILogSourceListProps> = props => {
    const { classes } = props;

    const [ storeState, store ] = useStoreStateFromContainerContext<ILogSourcePageStore, ILogSourcePageStoreState>(props);
    // // const  = useCallback(() => store..trigger(1), [ store ]);

    if (!storeState)
    {
        return null;
    }

    return <List className={classes.root}>
        { storeState.logSources && storeState.logSources.map(ls => {
            return <ListItem key={ls.id}>
                <ListItemText primary={ls.name + " " + ls.type}
                              secondary={ls.sourceFolder + " -> " + ls.fileRegex}></ListItemText>
            </ListItem>;
        }) }        
    </List>;
};

/**
 * Component that shows a list of log sources.
 */
export const LogSourceList = withStyles(styles)(LogSourceListImpl);
