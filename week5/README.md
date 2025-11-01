# Week 5

## Overview about Headless CMS

### What is Headless CMS?

Headless CMS is a backend-only content repository, where you can integrate any
type of frontend or UI libraries with. It's considered a cost-effective
solution for managing content, while allowing many ways of remixing through out
multi-media or digital channels.

### Comparing to traditional CMS

The traditional content management systems we knew before were using a type of
architecture that both the UI and the content are constrained by each other.
Therefore, information and code are tightly coupled (hence, coupled CMS as they
are now called), distributing the same content on a different media channel gets
complicated.

That's why headless CMS boasts the following merits:

- **Prioritized UX**: It's good for developers, content creators (not those
  "HEY GUYS WHAT'S GOING ON") and users all around. Each one can use one of the
  Headless CMS features: content creators can utilize a fully customized editor
  and a structural data model; developers can use filtering options, searches and
  queries to distribute content; and users can consume the content from whatever
  platform or frontend they are fond of.
- **Pretty nice integration with third parties**: You can manage who has access
  to your content. You can authorize a sharing with just a few clicks.

## Use cases for Headless CMS

It's that:

- ECommerce: you can personalize your items across multiple channels.
- Integration with other applications: there might be times when you need to
  share the content with not end users, but other software components like a
  chatbot, voice assistant, IoT devices or AI applications.
- Collaboration: regardless of frontend interfaces, you can manage the content
  from anywhere, any device.

## Headless CMS Architecture

Three main parts:

1. Content Repository: the place where you manage contents, databases and file
   systems.
2. API: The interface for acquiring, accessing or saving content with the
   repository. GraphQL, RestAPI with CRUD operations are sometimes provided.
3. Frontend: For system users, using the API to display the content to the user.
   This component is completely independent from the Headless CMS, developers may
   use any technology to implement this layer.

## Headless CMS Demonstration

Specifications:

- Software: **Directus**
- License: **BSL 1.1**
- Orchestrator: Docker Compose (recommended by Directus)

### Create a new collection

1. Login as admin account on Directus.
2. Create a new collection (SQLite database)

### CRUD on content

After creating a collection, content may be managed on a tab on the left toolbar.
It will tell you to fill in fields as you have created from the collection.

### Fetch content through an API

I use REST API to fetch the content through Directus, for example, to fetch the
collections `posts`, we can use:

```nu
http get https://directus-url/items/posts
```

Result:

```plaintext
╭──────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│      │ ╭───┬────┬──────────────────────────┬──────────────┬──────────────────────────────────┬─────────────────────────┬────────╮ │
│ data │ │ # │ id │       date_created       │ date_updated │              title               │         content         │ author │ │
│      │ ├───┼────┼──────────────────────────┼──────────────┼──────────────────────────────────┼─────────────────────────┼────────┤ │
│      │ │ 0 │  1 │ 2025-10-30T09:27:58.293Z │              │ Local woman bought Sakura's wand │ <p>Waow</p>             │      1 │ │
│      │ │ 1 │  2 │ 2025-10-30T09:29:04.566Z │              │ Blue shell breaks a man's kart   │ <p>He is very livid</p> │      3 │ │
│      │ ╰───┴────┴──────────────────────────┴──────────────┴──────────────────────────────────┴─────────────────────────┴────────╯ │
╰──────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
```

## References

- [What is Headless CMS? - AWS](https://aws.amazon.com/what-is/headless-cms/#:~:text=A%20headless%20content%20management%20system,effective%20solution%20for%20managing%20content.)
- [Directus Guides](https://directus.io/docs/guides/connect/query-parameters#page)
