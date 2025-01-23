import { Component, OnInit } from '@angular/core';
// import { ToolbarItems } from '@syncfusion/ej2-angular-grids';
// import { ToolbarModule } from '@syncfusion/ej2-angular-navigations';
// import { GridModule } from '@syncfusion/ej2-angular-grids';


@Component({
  selector: 'app-policies',
  templateUrl: './policies.component.html',
  styleUrls: ['./policies.component.css'],
  // imports: [ToolbarModule, GridModule],
})
export class PoliciesComponent implements OnInit {
  public policies: Policy[] = [];
  // public toolbarItems: ToolbarItems[] = ['Add', 'Edit', 'Delete', 'Search'];

  constructor() {
    // Sample data - replace with actual API call
    this.policies = [
      {
        policyId: 'POL001',
        policyName: 'Health Insurance',
        description: 'Basic health coverage policy',
        status: 'Active',
        createdDate: new Date('2024-01-15')
      },
      {
        policyId: 'POL002',
        policyName: 'Life Insurance',
        description: 'Term life insurance policy',
        status: 'Active',
        createdDate: new Date('2024-02-01')
      }
    ];
  }

  ngOnInit(): void {
  }
}

// Policy interface
interface Policy {
  policyId: string;
  policyName: string;
  description: string;
  status: string;
  createdDate: Date;
}
