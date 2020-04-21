import { handleAction, handleActionVoid, IInjectedStoreOptions, IStore, reduceAction, useStore } from "rfluxx";
import { clearMultiSlashes } from 'rfluxx-routing';

export interface ILogSource {
    id: number;
	name: string;
	type: string;
	sourceFolder: string;
	fileRegex: string;
}

/**
 * The state of the store { @see LogSourcePageStore }.
 */
export interface ILogSourcePageStoreState
{
    /**
     * A list of log sources to show.
     */
    logSources: ILogSource[];
}

/**
 * The options to configure the store { @see LogSourcePageStore }.
 */
export interface ILogSourcePageStoreOptions
    extends IInjectedStoreOptions
{
}

/**
 * The interface that exposes the commands of the store { @see LogSourcePageStore }.
 */
export interface ILogSourcePageStore extends IStore<ILogSourcePageStoreState>
{
    /**
     * Load the log sources from the backend.
     */
    loadLogSources();
}

/**
 * Store that manages the log source page.
 */
export const LogSourcePageStore = (options: ILogSourcePageStoreOptions) => {
    const initialState = {
        logSources: [
            {
                id: 1,
                name: "Log Source 1",
                type: "Job",
                sourceFolder: "/mnt/data/logs",
                fileRegex: "*.log"
            },
            {
                id: 2,
                name: "Log Source 2",
                type: "Job",
                sourceFolder: "/mnt/data/logs2",
                fileRegex: "*.log"
            }
        ]
    };
    const [state, setState, store] = useStore<ILogSourcePageStoreState>(initialState);

    var me = {
        ...store,
        loadLogSources: () => {
            fetch("/api/sources")
            .then(r => r.json())
            .then(json => {
                const logSources = json as ILogSource[];
                setState({
                    ...state.value,
                    logSources
                });
            })
        }
    }

    me.loadLogSources();

    return me;
};
