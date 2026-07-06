import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

/**
 * Merges Tailwind classes with proper precedence handling.
 * Uses clsx for conditional classes and tailwind-merge to resolve conflicts.
 */
export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

/**
 * Creates a URL-friendly slug from a string
 */
export function slugify(str) {
  return str
    .toLowerCase()
    .trim()
    .replace(/[^\w\s-]/g, "")
    .replace(/[\s_-]+/g, "-")
    .replace(/^-+|-+$/g, "");
}

/**
 * Deep clones an object
 */
export function deepClone(obj) {
  return JSON.parse(JSON.stringify(obj));
}

/**
 * Checks if value is empty (null, undefined, empty string, empty array, empty object)
 */
export function isEmpty(value) {
  if (value === null || value === undefined) return true;
  if (typeof value === "string") return value.trim() === "";
  if (Array.isArray(value)) return value.length === 0;
  if (typeof value === "object") return Object.keys(value).length === 0;
  return false;
}

/**
 * Capitalizes first letter of each word
 */
export function capitalize(str) {
  if (!str) return "";
  return str.replace(/\b\w/g, (char) => char.toUpperCase());
}

/**
 * Truncates text with ellipsis
 */
export function truncate(str, length = 50) {
  if (!str || str.length <= length) return str;
  return str.slice(0, length) + "...";
}

/**
 * Calculates elapsed minutes between two dates
 */
export function getElapsedMinutes(dateString) {
  const start = new Date(dateString);
  const now = new Date();
  return Math.floor((now - start) / (1000 * 60));
}

/**
 * Gets hold bill color config based on elapsed minutes
 */
export function getHoldBillColor(elapsedMinutes) {
  if (elapsedMinutes > 60) {
    return {
      border: "border-danger-500",
      bg: "bg-danger-50",
      text: "text-danger-600",
    };
  }
  if (elapsedMinutes > 30) {
    return {
      border: "border-secondary-500",
      bg: "bg-secondary-50",
      text: "text-secondary-600",
    };
  }
  return {
    border: "border-primary-500",
    bg: "bg-primary-50",
    text: "text-primary-600",
  };
}

/**
 * Validates Indonesian phone number
 */
export function isValidIndonesianPhone(phone) {
  return /^08\d{8,11}$/.test(phone);
}

/**
 * Generates a unique ID
 */
export function generateId(prefix = "id") {
  return `${prefix}_${Math.random().toString(36).substr(2, 9)}_${Date.now().toString(36)}`;
}

/**
 * Safely parses JSON
 */
export function safeJsonParse(str, fallback = null) {
  try {
    return JSON.parse(str);
  } catch {
    return fallback;
  }
}

/**
 * Groups array by key
 */
export function groupBy(array, key) {
  return array.reduce((result, item) => {
    const group = item[key];
    result[group] = result[group] || [];
    result[group].push(item);
    return result;
  }, {});
}

/**
 * Sorts array of objects by key
 */
export function sortBy(array, key, direction = "asc") {
  return [...array].sort((a, b) => {
    const aVal = a[key];
    const bVal = b[key];
    if (aVal < bVal) return direction === "asc" ? -1 : 1;
    if (aVal > bVal) return direction === "asc" ? 1 : -1;
    return 0;
  });
}

/**
 * Calculates percentage change
 */
export function calculatePercentChange(current, previous) {
  if (!previous || previous === 0) return 0;
  return Math.round(((current - previous) / previous) * 100);
}

/**
 * Formats file size
 */
export function formatFileSize(bytes) {
  if (bytes === 0) return "0 B";
  const k = 1024;
  const sizes = ["B", "KB", "MB", "GB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
}
