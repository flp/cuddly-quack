# Roadmap

This is a proposed Roadmap for our MTG draft simulation project.

## Featureset

This is the initial featureset for a bare-minimum drafting experience (read: we are not trying to create a full play experience just yet):

- [ ] Drafts
  - [ ] Players may create 3-round swiss drafts of the current set
    - [ ] Draft rooms can be accessed by UUID
    - [ ] Draft rooms require a password by default
    - [ ] [optional] Public draft rooms

- [ ] Players 
  - [ ] Players are anonymous by default
  - [ ] [optional] Allow players to sign in via an Oauth token
    - [ ] FB / Twitter / Google / w/e
  - [ ] Players can save their draft decks in a deck list for later
    - [ ] export as `.dek` so we can use it for other programs, like XMage, or w/e

- [ ] Cards
  - [ ] Pull images dynamically from mythicspolier.com at first as to not go against Wizard's wishes

- [ ] Drafting
  - [ ] Create "Packs" consisting of 1 rare, 3 uncommons, 9 commons, and a land
    - [ ] Provide mythic support
    - [ ] Provide foil support
  - [ ] Issue 3 "Packs" randomly to all members of the draft
  - [ ] Exchange packs player by player in Left, Right, Left format
    - [ ] Implement a timer so that players are forced to choose within the apropriate time limit
    - [ ] Implement an auto choose function when they run out of time
    - [ ] Pass remainder of the pack to the person next in line