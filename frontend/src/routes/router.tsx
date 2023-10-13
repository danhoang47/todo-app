import {
	createBrowserRouter,
	createRoutesFromElements,
	Route,
} from "react-router-dom";

import App from "@/App";
import { Home } from "@/pages";

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route path="/" element={<App />}>
            <Route element={<Home />} >
                <Route index element={<div>Today</div>} />
                <Route path="?upcoming" element={<div>upcoming</div>} />
            </Route>
        </Route>
    )
);

export default router;
