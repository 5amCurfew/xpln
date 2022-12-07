# Load dplyr library
library('dplyr')
df %>% 
    filter(state %in% c("CA", "AZ", "PH")) %>% 
    group_by(gender) %>%
    summarise(total = n())