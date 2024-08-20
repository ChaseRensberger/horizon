
channels
 - channel_id (PK)
 - channel_name
 - country
 - custom_url

channel_snapshots
 - channel_snapshot_id (PK)
 - channel_id (FK -> channels)
 - view_count
 - subscriber_count
 - video_count
 - hidden_subscriber_count

videos
 - video_id (PK)
 - channel_id (FK -> channels)

video_snapshots
 - video_snapshot_id
 - video_id (FK -> videos)
 - view_count
 - like_count
 - comment_count 
  



## random thoughts
- I assume snapshots are important? also wtf is `favoriteCount`

## stuff to do
