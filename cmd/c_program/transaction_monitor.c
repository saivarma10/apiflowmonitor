#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>

struct ResponseData {
    char *data;
    size_t size;
};

size_t write_callback(void *ptr, size_t size, size_t nmemb, struct ResponseData *response) {
    size_t new_size = response->size + (size * nmemb);
    response->data = realloc(response->data, new_size + 1);
    
    if (response->data == NULL) {
        fprintf(stderr, "Failed to allocate memory for response data\n");
        return 0;
    }
    
    memcpy(response->data + response->size, ptr, size * nmemb);
    response->data[new_size] = '\0';
    response->size = new_size;
    
    return size * nmemb;
}

struct curl_slist* set_headers(const char *auth_token, const char *extra_headers) {
    struct curl_slist *headers = NULL;

    headers = curl_slist_append(headers, "Content-Type: application/json");

    if (auth_token != NULL && strlen(auth_token) > 0) {
        char auth_header[512];
        snprintf(auth_header, sizeof(auth_header), "Authorization: Bearer %s", auth_token);
        headers = curl_slist_append(headers, auth_header);
    }

    if (extra_headers != NULL && strlen(extra_headers) > 0) {
        char *header = strdup(extra_headers);
        if (header == NULL) {
            fprintf(stderr, "Memory allocation failed for headers\n");
            curl_slist_free_all(headers);
            return NULL;
        }

        char *token = strtok(header, ",");
        while (token != NULL) {
            headers = curl_slist_append(headers, token);
            token = strtok(NULL, ",");
        }
        free(header);
    }

    return headers;
}

void make_request(const char *url, const char *method, const char *headers, 
                 const char *auth_token, const char *payload) {
    CURL *curl;
    CURLcode res;
    long response_code;
    double response_time;
    struct ResponseData response = {
        .data = malloc(1),
        .size = 0
    };
    
    if (response.data == NULL) {
        fprintf(stderr, "Initial memory allocation failed\n");
        return;
    }
    response.data[0] = '\0';

    if (url == NULL || method == NULL) {
        fprintf(stderr, "URL and method cannot be NULL.\n");
        free(response.data);
        return;
    }

    curl_global_init(CURL_GLOBAL_DEFAULT);
    curl = curl_easy_init();

    if (curl) {
        struct curl_slist *curl_headers = set_headers(auth_token, headers);
        if (headers != NULL && curl_headers == NULL) {
            fprintf(stderr, "Failed to set headers.\n");
            free(response.data);
            curl_easy_cleanup(curl);
            curl_global_cleanup();
            return;
        }

        curl_easy_setopt(curl, CURLOPT_URL, url);

        if (curl_headers != NULL) {
            curl_easy_setopt(curl, CURLOPT_HTTPHEADER, curl_headers);
        }

        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, 1L);
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, 2L);

        if (strcmp(method, "GET") == 0) {
            curl_easy_setopt(curl, CURLOPT_HTTPGET, 1L);
        } else if (strcmp(method, "POST") == 0) {
            curl_easy_setopt(curl, CURLOPT_POST, 1L);
            if (payload != NULL && strlen(payload) > 0) {
                curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
            }
        } else if (strcmp(method, "PUT") == 0) {
            curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "PUT");
            if (payload != NULL && strlen(payload) > 0) {
                curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
            }
        } else if (strcmp(method, "DELETE") == 0) {
            curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "DELETE");
        } else {
            fprintf(stderr, "Unsupported HTTP method: %s\n", method);
            free(response.data);
            curl_easy_cleanup(curl);
            curl_slist_free_all(curl_headers);
            curl_global_cleanup();
            return;
        }

        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);

        curl_easy_setopt(curl, CURLOPT_TIMEOUT, 30L);
        curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 10L);

        res = curl_easy_perform(curl);

        if (res != CURLE_OK) {
            fprintf(stderr, "Request failed: %s\n", curl_easy_strerror(res));
        } else {
            curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &response_code);
            curl_easy_getinfo(curl, CURLINFO_TOTAL_TIME, &response_time);

            printf("{\n"
                   "  \"status_code\": %ld,\n"
                   "  \"response_time\": %.3f,\n"
                   "  \"response_data\": %s\n"
                   "}\n",
                   response_code, response_time, 
                   (response.size > 0) ? response.data : "\"\"");
        }

        curl_easy_cleanup(curl);
        if (curl_headers) {
            curl_slist_free_all(curl_headers);
        }
    }

    free(response.data);
    curl_global_cleanup();
}

int main(int argc, char *argv[]) {
    if (argc < 3) {
        fprintf(stderr, "Usage: %s <url> <method> [auth_token] [headers] [payload]\n", argv[0]);
        return 1;
    }

    const char *url = argv[1];
    const char *method = argv[2];
    const char *auth_token = (argc > 3) ? argv[3] : NULL;
    const char *headers = (argc > 4) ? argv[4] : NULL;
    const char *payload = (argc > 5) ? argv[5] : NULL;

    make_request(url, method, headers, auth_token, payload);
    return 0;
}