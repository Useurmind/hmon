import { handleAction, handleActionVoid, IInjectedStoreOptions, IStore, reduceAction, useStore } from "rfluxx";
import { RouteParameters } from 'rfluxx-routing';
import { Observable } from "rxjs";

import { ILogSource } from "../log_sources/LogSourcePageStore";


export interface IJobLogSource extends ILogSource {
    successRegex: string;
    errorRegex: string;
}

/**
 * The state of the store { @see JobLogSourceStore }.
 */
export interface IJobLogSourceStoreState
{
    /**
     * The id of the job log source to manage (or null if new).
     */
    jobLogSourceId: number | null;

    /**
     * A string value set by the example action @setString.
     */
    jobLogSource: IJobLogSource;
}

/**
 * The options to configure the store { @see JobLogSourceStore }.
 */
export interface IJobLogSourceStoreOptions
    extends IInjectedStoreOptions
{
    /**
     * Route parameters of the page.
     */
    routeParameters: Observable<RouteParameters>
}

/**
 * The interface that exposes the commands of the store { @see JobLogSourceStore }.
 */
export interface IJobLogSourceStore extends IStore<IJobLogSourceStoreState>
{
    /**
     * Load the job log source with the given id.
     */
    load(jobLogSourceId: number);

    /**
     * Save the current job log source.
     */
    save();

    /**
     * Set the name of the jls
     * @param name The name.
     */
    setName(name: string);

    setSourceFolder(sourceFolder: string);
    setFileRegex(fileRegex: string);
    setSuccessRegex(value: string);
    setErrorRegex(value: string);
}

/**
 * Store that manages a single job log source.
 */
export const JobLogSourceStore = (options: IJobLogSourceStoreOptions) => {
    const initialState = {
        jobLogSourceId: null,
        jobLogSource: {
            id: null,
            type: "job"
        }
    } as IJobLogSourceStoreState;
    const [state, setState, store] = useStore<IJobLogSourceStoreState>(initialState);

    options.routeParameters.subscribe(p =>
    {
        const isCreate = p.get("id") === "new";
        const jlsId = isCreate ? -1 : p.getAsInt("id");

        if (!isCreate)
        {
            fetch(`/api/jobsource/${jlsId}`)
                .then(r => r.json())
                .then(json =>
                {
                    const jobLogSource = json as IJobLogSource;

                    setState({
                        ...state.value,
                        jobLogSourceId: jobLogSource.id,
                        jobLogSource
                    })
                })
        }
    });

    return {
        ...store,
        save: () =>
        {
            fetch("/api/jobsource", {
                method: "POST",
                body: JSON.stringify(state.value.jobLogSource),
                cache: "no-cache",
                headers: {
                    "Content-Type": "application/json"
                }
            })
            .then(r => r.json())
            .then(json =>
            {
                const newLogSource = json as IJobLogSource;

                setState({
                    ...state.value,
                    jobLogSourceId: newLogSource.id,
                    jobLogSource: newLogSource
                })
            })
        },
        setName:reduceAction(state, (s, name: string) =>
        {
            return { ...s, jobLogSource: { ...s.jobLogSource, name }}
        }),
        setSourceFolder:reduceAction(state, (s, sourceFolder: string) =>
        {
            return { ...s, jobLogSource: { ...s.jobLogSource, sourceFolder }}
        }),
        setFileRegex:reduceAction(state, (s, fileRegex: string) =>
        {
            return { ...s, jobLogSource: { ...s.jobLogSource, fileRegex }}
        }),
        setSuccessRegex:reduceAction(state, (s, successRegex: string) =>
        {
            return { ...s, jobLogSource: { ...s.jobLogSource, successRegex }}
        }),
        setErrorRegex:reduceAction(state, (s, errorRegex: string) =>
        {
            return { ...s, jobLogSource: { ...s.jobLogSource, errorRegex }}
        })
    }
};
