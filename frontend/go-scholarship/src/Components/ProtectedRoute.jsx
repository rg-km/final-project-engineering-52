import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../Database/useAuth";

export function RequireAuth({ children }) {
  let { user } = useAuth((state) => state);
  let location = useLocation();

  if (JSON.stringify(user) === '{}') {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
}
