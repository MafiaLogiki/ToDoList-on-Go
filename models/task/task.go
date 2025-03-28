package task

type Task struct {
    Id            int    `json:"id"`
    Title         string `json:"title"`
    Description   string `json:"descriprion"`
    Status        string `json:"status"`
    UserId        int    `json:"userId"`
}
