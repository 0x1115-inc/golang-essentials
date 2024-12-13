# Message System

## Google Cloud Pub/Sub setup
1. Create a new project in Google Cloud Console.
2. Enable Pub/Sub API.
3. Create a new topic.
```bash
PUBSUB_TOPIC_NAME="my-topic"

gcloud pubsub topics create $PUBSUB_TOPIC_NAME
```
4. Create a new subscription.
```bash
PUBSUB_SUBSCRIPTION_NAME="my-subscription"

gcloud pubsub subscriptions create $PUBSUB_SUBSCRIPTION_NAME --topic=$PUBSUB_TOPIC_NAME
```
5. Create a service account with Pub/Sub Publisher and Subscriber roles.
```bash
SERVICE_ACCOUNT_NAME="my-service-account"
SERVICE_ACCOUNT_EMAIL="$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com"

# Role for Publisher application
ROLES="roles/pubsub.publisher"

# Role for Subscriber application
ROLES="roles/pubsub.subscriber"

# Attach roles to the service account
gcloud projects add-iam-policy-binding $PROJECT_ID --member=serviceAccount:$SERVICE_ACCOUNT_EMAIL --role=$ROLES
```

## References
- A great inspiration came from [Enterprise Integration Patterns blog](https://l.0x1115.com/Cz8B2KtU9dc).